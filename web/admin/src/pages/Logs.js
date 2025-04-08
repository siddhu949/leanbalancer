import React, { useEffect } from "react";
import gsap from "gsap";
import "../styles.css"; // Ensure correct import path

const handleHover = (e) => {
  gsap.to(e.target, {
    backgroundColor: "#ffcc00",
    color: "#000",
    scale: 1.05,
    duration: 0.3,
  });
};

const handleHoverExit = (e) => {
  gsap.to(e.target, {
    backgroundColor: "#333",
    color: "#fff",
    scale: 1,
    duration: 0.3,
  });
};

const Logs = () => {
  useEffect(() => {
    gsap.from(".logs", {
      opacity: 0,
      y: 30,
      duration: 1,
      ease: "power3.out",
    });
  }, []);

  return (
    <div className="logs">
      <h2 className="logs-title">📜 Logs</h2>
      <div className="log-entry" onMouseEnter={handleHover} onMouseLeave={handleHoverExit}>
        Log data will appear here...
      </div>
      <div className="log-entry" onMouseEnter={handleHover} onMouseLeave={handleHoverExit}>
        Example log entry 1...
      </div>
      <div className="log-entry" onMouseEnter={handleHover} onMouseLeave={handleHoverExit}>
        Example log entry 2...
      </div>
    </div>
  );
};

export default Logs;
