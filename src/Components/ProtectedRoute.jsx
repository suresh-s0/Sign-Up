


import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
  // const users = JSON.parse(localStorage.getItem("user")) || [];
  const isAuthenticated = localStorage.getItem("isAuthenticated") === "true";

  if (!isAuthenticated ) {
    return <Navigate to="/" />;
  }
  return children;
};

export default ProtectedRoute;
