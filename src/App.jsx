import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import SignIn from "./Components/SignIn";
import SignUp from "./Components/Signup";    

export default function App() {
  return (
    <div className="min-h-screen bg-gray-900">
      <Router>
      
        <Routes>
          <Route path="/" element={<SignUp/>} />  
          <Route path="/SignIn" element={<SignIn/>} />  
        </Routes>
      </Router>
    </div>
  );
}
