import { Navigate } from "react-router-dom";
import * as jwtDecode from "jwt-decode";

const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem("token");

  const isAuthenticated = token && isTokenValid(token);

  if (!isAuthenticated) {
    return <Navigate to="/signin" />;
  }

  return children;
};

const isTokenValid = (token) => {
  try {
    const decodedToken = jwtDecode(token);

    if (decodedToken.exp * 1000 < Date.now()) {
      return false;
    }

    return true;
  } catch (error) {
    return false;
  }
};

export default ProtectedRoute;
