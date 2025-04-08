import React, { useEffect } from "react";
import gsap from "gsap";
import "./styles.css";

const LoadingScreen = ({ setLoading }) => {
  useEffect(() => {
    // Admin logo fade-in
    gsap.fromTo(
      ".loading-logo",
      { opacity: 0, y: -20 },
      { opacity: 1, y: 0, duration: 1, ease: "power2.out" }
    );

    // Server animation - bounce in sequence
    gsap.fromTo(
      ".server",
      { opacity: 0, y: 20 },
      {
        opacity: 1,
        y: 0,
        duration: 0.8,
        stagger: 0.2,
        ease: "bounce.out",
      }
    );

    // Glowing effect for active servers
    gsap.to(".server", {
      boxShadow: "0 0 15px rgba(0, 162, 255, 0.8)",
      repeat: -1,
      yoyo: true,
      duration: 1.5,
    });

    // Loading text animation
    gsap.to(".loading-text", {
      opacity: 0.5,
      repeat: -1,
      yoyo: true,
      duration: 1,
    });

    // Simulate loading completion
    setTimeout(() => setLoading(false), 4000);
  }, [setLoading]);

  return (
    <div className="loading-screen">
      <img src="/admin-logo.png" alt="Admin Logo" className="loading-logo" />
      <h2 className="loading-title">Memory Efficient Load Balancer</h2>
      <div className="server-container">
        <div className="server"></div>
        <div className="server"></div>
        <div className="server"></div>
      </div>
      <p className="loading-text">Loading...</p>
    </div>
  );
};

export default LoadingScreen;
