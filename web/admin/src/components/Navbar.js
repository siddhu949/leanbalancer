import React, { useState, useEffect } from "react";
import "../styles.css";

const Navbar = () => {
  const [time, setTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => setTime(new Date()), 1000);
    return () => clearInterval(timer);
  }, []);

  return (
    <nav className="navbar">
      {/* Admin Logo */}
      <div className="admin-logo">
        <img src="/admin-logo.png" alt="Admin Logo" />
      </div>

      {/* Title with Gradient Color */}
      <h1 className="navbar-title">Admin Dashboard</h1>

      {/* Time with Gradient Color */}
      <span className="clock">{time.toLocaleTimeString()}</span>
    </nav>
  );
};

export default Navbar;
