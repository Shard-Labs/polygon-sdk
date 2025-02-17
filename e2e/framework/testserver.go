package framework

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/0xPolygon/minimal/minimal/proto"
	"github.com/0xPolygon/minimal/types"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/umbracle/go-web3/jsonrpc"
	"google.golang.org/grpc"
)

type TestServerConfig struct {
	JsonRPCPort  int64
	GRPCPort     int64
	LibP2PPort   int64
	Seal         bool
	DataDir      string
	PremineAccts map[types.Address]*big.Int
}

func (t *TestServerConfig) Premine(addr types.Address, amount *big.Int) {
	if t.PremineAccts == nil {
		t.PremineAccts = map[types.Address]*big.Int{}
	}
	t.PremineAccts[addr] = amount
}

type TestServerConfigCallback func(*TestServerConfig)

var initialPort = int64(12000)

func getOpenPort() int64 {
	port := atomic.AddInt64(&initialPort, 1)
	return port
}

type TestServer struct {
	t *testing.T

	config *TestServerConfig
	cmd    *exec.Cmd
}

func (t *TestServer) GrpcAddr() string {
	return fmt.Sprintf("http://127.0.0.1:%d", t.config.GRPCPort)
}

func (t *TestServer) JsonRPCAddr() string {
	return fmt.Sprintf("http://127.0.0.1:%d", t.config.JsonRPCPort)
}

func (t *TestServer) JSONRPC() *jsonrpc.Client {
	fmt.Println("////")
	fmt.Println(t.JsonRPCAddr())

	clt, err := jsonrpc.NewClient(t.JsonRPCAddr())
	if err != nil {
		t.t.Fatal(err)
	}
	return clt
}

func (t *TestServer) Operator() proto.SystemClient {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", t.config.GRPCPort), grpc.WithInsecure())
	if err != nil {
		t.t.Fatal(err)
	}
	return proto.NewSystemClient(conn)
}

func (t *TestServer) Stop() {
	if err := t.cmd.Process.Kill(); err != nil {
		t.t.Error(err)
	}
}

func NewTestServer(t *testing.T, callback TestServerConfigCallback) *TestServer {
	path := "polygon-sdk"

	config := &TestServerConfig{
		JsonRPCPort: getOpenPort(),
		GRPCPort:    getOpenPort(),
		LibP2PPort:  getOpenPort(),
	}

	dataDir, err := ioutil.TempDir("/tmp", "polygon-sdk-e2e-")
	if err != nil {
		t.Fatal(err)
	}

	config.DataDir = dataDir
	if callback != nil {
		callback(config)
	}

	// Build genesis file
	{
		args := []string{
			"genesis",
			// add data dir
			"--data-dir", dataDir,
		}
		// add premines
		for addr, amount := range config.PremineAccts {
			args = append(args, "--premine", addr.String()+":0x"+amount.Text(16))
		}

		vcmd := exec.Command(path, args...)
		vcmd.Stdout = nil
		vcmd.Stderr = nil
		if err := vcmd.Run(); err != nil {
			t.Skipf("polygon-sdk genesis failed: %v", err)
		}
	}

	// Build arguments
	args := []string{
		"server",
		// add data dir
		"--data-dir", dataDir,
		// add custom chain
		"--chain", filepath.Join(dataDir, "genesis.json"),
		// enable grpc
		"--grpc", fmt.Sprintf(":%d", config.GRPCPort),
		// enable libp2p
		"--libp2p", fmt.Sprintf(":%d", config.LibP2PPort),
		// enable jsonrpc
		"--jsonrpc", fmt.Sprintf(":%d", config.JsonRPCPort),
	}

	if config.Seal {
		args = append(args, "--seal")
	}

	stdout := io.Writer(os.Stdout)
	stderr := io.Writer(os.Stdout)

	fmt.Println(strings.Join(args, " "))

	// Start the server
	cmd := exec.Command(path, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("err: %s", err)
	}

	srv := &TestServer{
		t:      t,
		config: config,
		cmd:    cmd,
	}

	// wait until is ready
	for {
		if _, err := srv.Operator().GetStatus(context.Background(), &empty.Empty{}); err == nil {
			break
		}
	}
	return srv
}
