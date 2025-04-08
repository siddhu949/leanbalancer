import React, { useState, useEffect, useRef } from "react";
import gsap from "gsap";
import { ReactComponent as SunIcon } from "../assets/sun.svg";
import { ReactComponent as MoonIcon } from "../assets/moon.svg";
import "../styles.css";

const ThemeToggle = () => {
  const [darkMode, setDarkMode] = useState(() => {
    return localStorage.getItem("theme") === "dark";
  });

  const toggleRef = useRef(null);
  const sunRef = useRef(null);
  const moonRef = useRef(null);

  useEffect(() => {
    document.body.classList.toggle("dark-mode", darkMode);
    localStorage.setItem("theme", darkMode ? "dark" : "light");

    if (darkMode) {
      gsap.to(moonRef.current, {
        opacity: 1,
        scale: 1,
        duration: 0.5,
        ease: "elastic.out(1, 0.5)",
      });
      gsap.to(sunRef.current, {
        opacity: 0,
        scale: 0,
        duration: 0.3,
        ease: "power3.inOut",
      });
    } else {
      gsap.to(sunRef.current, {
        opacity: 1,
        scale: 1,
        duration: 0.5,
        ease: "elastic.out(1, 0.5)",
      });
      gsap.to(moonRef.current, {
        opacity: 0,
        scale: 0,
        duration: 0.3,
        ease: "power3.inOut",
      });
    }
  }, [darkMode]);

  const handleToggle = () => {
    setDarkMode(!darkMode);
    gsap.fromTo(
      toggleRef.current,
      { scale: 0.9 },
      { scale: 1, duration: 0.2, ease: "bounce.out" }
    );
  };

  return (
    <div className="theme-toggle-container" onClick={handleToggle}>
      <div ref={toggleRef} className="theme-toggle">
        <SunIcon ref={sunRef} className="icon sun-icon" />
        <MoonIcon ref={moonRef} className="icon moon-icon" />
      </div>
    </div>
  );
};

export default ThemeToggle;
