module.exports = {
  compilers: {
    solc: {
      version: "0.8.4+commit.c7e474f2.Emscripten.clang"
    }
  },
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    },
    test: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    }
  }
};
