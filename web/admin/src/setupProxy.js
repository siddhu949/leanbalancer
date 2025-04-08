const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function(app) {
  app.use(
    "/metrics",
    createProxyMiddleware({
      target: "http://localhost:9002/metrics", // Your backend URL
      changeOrigin: true,
      secure: false, // If using HTTPS without a valid certificate
    })
  );
};
