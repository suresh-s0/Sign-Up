import {  useState } from "react";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function SignIn() {
  const [email, setEmail] = useState("");
  const [password, setPass] = useState("");
  const navigate = useNavigate();
  // const users = JSON.parse(localStorage.getItem("user")) || [];
  const handleclic = async (e) => {
    e.preventDefault();
    let newUSer = {
      email,
      password,
    };

    const token = localStorage.getItem("jwtToken");
    try {
      let response = await fetch("http://localhost:8080/api/signin", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
            "Authorization": token ? `Bearer ${token}` : ""
        },
        body: JSON.stringify(newUSer),
      });
      if (response.ok) {
        const data = await response.json();
        console.log(data)

        const token= data.token

        if (token) {
          // Save token in localStorage
          localStorage.setItem("jwtToken", token);
          localStorage.setItem("isAuthenticated", "true");

          toast.success("Login successful!", {
            position: "top-center",
          });
        }
        navigate("/Home");
      } else {
        // Failed login
        const errorData = await response.json(); 
        toast.error(errorData.message || "Email or password does not match", {
          position: "top-center",
        });
      }
    } catch (error) {
      console.error("Error:", error);
      console.error("Error:", error);
      toast.error(`Error: ${error.message}`, {
        position: "top-center",
      });
    }
  };
  

  return (
    <div className="flex flex-col  justify-center items-center     p-6">
      <div className="flex flex-col  justify-center items-center  m-[100px]  sm:w-full md:w-2/3 lg:w-1/2 xl:w-full   px-6 py-4 ">
        <h1 className="font-bold text-4xl text-white">Sign-In </h1>

        <div className="flex flex-col justify-center    sm:w-full md:w-1/2 lg:w-1/3 xl:w-1/3 m-4  p-10 border">
          <label htmlFor="email" className="text-white font-semibold">
            Email
          </label>
          <input
            className="border  rounded rounded:md p-2 mb-2"
            type="email"
            placeholder="Enter Your Email"
            onChange={(e) => setEmail(e.target.value)}
          />

          <label htmlFor="pasword" className="text-white font-semibold">
            Password
          </label>
          <input
            className="border  rounded-md p-2 mb-2"
            type="password"
            placeholder="Enter Your password"
            onChange={(e) => setPass(e.target.value)}
          />

          <button
            className=" text-white bg-green-800   flex items-center  rounded-lg   font-semibold border w-fit px-4 py-2 mx-auto"
            type="button"
            onClick={handleclic}
          >
            Submit
          </button>
        </div>
        <ToastContainer />
      </div>
    </div>
  );
}

export default SignIn;
