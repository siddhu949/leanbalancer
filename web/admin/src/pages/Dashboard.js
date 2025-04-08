import React from "react";
import "../styles.css";

const Dashboard = () => {
  return (
    <div className="dashboard-wrapper">
      {/* Admin Dashboard Section */}
      <section className="admin-dashboard">
        <h1>👨‍💼 Admin Dashboard</h1>
        <p>Welcome to the centralized monitoring system.</p>
        <div className="dashboard-cards">

          {/* Load Balancer Card - Landscape Format */}
          <div className="dashboard-card load-balancer-card">
           
            <div className="load-balancer-card-content">
              <h5>🔀Memory-Efficient Load Balancer</h5>

              <p>🔹 <strong>Features:</strong></p>
              <ul>
                <li>Forward & Reverse Proxy</li>
                <li>High Performance Routing</li>
                <li>Load Distribution</li>
                <li>Security Enhancements</li>
              </ul>

              <p>🎯 <strong>Objectives:</strong></p>
              <ul>
                <li>Efficient Resource Utilization</li>
                <li>Fault Tolerance</li>
                <li>Scalability</li>
              </ul>

            

              
            </div>
          </div>

        </div>
        <a href="#grafana" className="btn scroll-btn">View Grafana Dashboard</a>
      </section>

      {/* Grafana Dashboard Section */}
      <section id="grafana" className="grafana-dashboard">
        <h1>📈 Grafana Dashboard</h1>
        <p>Real-time monitoring with Grafana.</p>
        <iframe
          src="http://localhost:3000/goto/HCv_it0HR?orgId=1"
          width="100%"
          height="800px"
          style={{ border: "none", maxWidth: "100%" }}
          title="Grafana Dashboard"
        ></iframe>
      </section>

      {/* Additional Sections */}
      <section className="monitoring-section">
        <h2>📊 Prometheus Monitoring</h2>
        <p>Collect and store time-series data with PromQL.</p>
        <a href="http://localhost:9090" target="_blank" rel="noopener noreferrer" className="btn">
          Open Prometheus
        </a>
      </section>

      <section className="monitoring-section">
        <h2>⚡ Load Balancer Optimization</h2>
        <p>Smart traffic routing for high availability and efficiency.</p>
        <a href="https://github.com/siddhu949/leanbalancer.git" target="_blank" rel="noopener noreferrer" className="btn">
          Learn More
        </a>
      </section>
    </div>
  );
};

export default Dashboard;
