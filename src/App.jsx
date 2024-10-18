import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import SignIn from "./Components/SignIn";
import SignUp from "./Components/Signup";  
import Home from "./Components/Home"; 
import ProtectedRoute from './Components/ProtectedRoute';
import { ToastContainer } from "react-toastify";

export default function App() {
  return (
    <div className="min-h-screen bg-gray-900">
        <ToastContainer />
      <Router>
    
        <Routes>
          <Route path="/" element={<SignUp/>} />  
          <Route path="/SignIn" element={<SignIn/>} />  
          <Route path="/Home" element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          } />  

        </Routes>
      </Router>
    </div>
  );
}
