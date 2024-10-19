import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function SignUp() {
  const [users, setUser] = useState([]);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPass] = useState("");
  const navigate = useNavigate();

  const handleclic = async (e) => {
    e.preventDefault();

    if (!name || !email || !password) {
      toast.error("empty fields are not allowed", {
        position: "top-center",
      });
      return;
    }

    let newUser = {
      name,
      email,
      password,
    };
    try {
      const response = await fetch("http://localhost:8080/api/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newUser),
      });
      if (!response.ok) {
        const errorData = await response.json();
        toast.error(errorData.message || "An error occurred");
        return;
      }
      const data = await response.json();
      console.log(data);
      const token = data.data;  
      console.log(token);


      if (token) {
        let updatedUser = [...users, newUser];
      setUser(updatedUser);
      localStorage.setItem("users", JSON.stringify(updatedUser));
     
        localStorage.setItem("jwtToken", token);

        localStorage.setItem("isAuthenticated", "true");
  
        toast.success("Signup successful!", {
          position: "top-center",
        });
  
        navigate("/signin");  // Navigate to signin page after successful signup
      }






    } catch (error) {
      console.error("Error:", error);
      console.error("Error:", error);
      toast.error(`Error: ${error.message}`, {
        position: "top-center",
      });
    }
  };


  const handleSignIn = () => {
    navigate("/SignIn");
  };

  // to keep the localstorage
  useEffect(() => {
    let savedUsers = JSON.parse(localStorage.getItem("users"));
    if (savedUsers) {
      setUser(savedUsers);
    }
  }, []);

  return (
    <div className="flex flex-col  justify-center items-center font-monospace p-6 ">
      <div className="flex flex-col  justify-center items-center m-[100px]  sm:w-full md:w-2/3 lg:w-full xl:w-full   px-6 py-4 ">
        <h1 className="font-bold text-4xl text-white">Sign-Up </h1>

        <div className="flex flex-col justify-center  p-10  sm:w-full md:w-1/full lg:w-1/3 xl:w-1/3 m-4  border">
          <label htmlFor="name" className="text-white font-semibold">
            Name
          </label>
          <input
            className="border  rounded rounded:lg p-2 mb-2"
            type="text"
            placeholder="Enter Your name "
            value={name}
            onChange={(e) => setName(e.target.value)}
          />

          <label htmlFor="email" className="text-white font-semibold">
            Email
          </label>
          <input
            className="border  rounded rounded:md p-2 mb-2"
            type="email"
            placeholder="Enter Your Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <label htmlFor="pasword" className="text-white font-semibold">
            Password
          </label>
          <input
            className="border  rounded-md p-2 mb-2"
            type="password"
            placeholder="Enter Your password"
            value={password}
            onChange={(e) => setPass(e.target.value)}
          />

          <button
            className=" text-white bg-green-800   flex items-center  rounded-lg   font-semibold border w-fit px-4 py-2 mx-auto"
            type="button"
            onClick={handleclic}
          >
            Submit
          </button>

          <button
            className=" text-white bg-green-800   flex items-center rounded-lg   font-semibold border w-fit px-3 py-2 mt-2 mx-auto"
            type="button"
            onClick={handleSignIn}
          >
            Sign-In
          </button>
        </div>
      </div>
  
    </div>
  );
}
export default SignUp;
