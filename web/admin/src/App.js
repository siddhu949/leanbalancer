import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import gsap from "gsap";
import LoadingScreen from "./LoadingScreen";
import Navbar from "./components/Navbar";
import Dashboard from "./pages/Dashboard";
import "./styles.css";

const App = () => {
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!loading) {
      gsap.set(".app-container", { opacity: 0, y: 50 });
      gsap.to(".app-container", { opacity: 1, y: 0, duration: 1 });
    }
  }, [loading]);

  return (
    <Router>
      {loading ? (
        <LoadingScreen setLoading={setLoading} />
      ) : (
        <div className="app-container">
          <Navbar />
          <div className="content">
            <Routes>
              <Route path="/" element={<Dashboard />} />
            </Routes>
          </div>
        </div>
      )}
    </Router>
  );
};

export default App;
