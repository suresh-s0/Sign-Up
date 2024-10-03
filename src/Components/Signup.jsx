import { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';

function SignUp() {
  const [users, setUser] = useState([]);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPass] = useState("");

  const handleclic = (e) => {
    e.preventDefault();
    if (!name || !email || !password) {
      window.alert("empty fields");
      return;
    }
    let newUser = {
      name,
      email,
      password,
    };

    let User = [...users, newUser];

    setUser(User);
    localStorage.setItem("user", JSON.stringify(User));
  };
   
  const navigate= useNavigate();

  const handleSignIn=()=>{
    
    navigate("/SignIn")

  }

  // to keep the localstorage
  useEffect(() => {
    let savedUsers = JSON.parse(localStorage.getItem("user"));
    if (savedUsers) {
      setUser(savedUsers);
    }
  }, []);

  return (
    <div className="flex flex-col  justify-center items-center  border-2 p-6 ">
      <div className="flex flex-col  justify-center items-center m-[100px] sm:w-full md:w-2/3 lg:w-1/2 xl:w-full border-2  px-6 py-4 ">
        <h1 className="font-bold text-4xl text-white">Sign-Up </h1>

        <div className="flex flex-col justify-center  p-3  sm:w-full md:w-1/3 lg:w-1/2 xl:w-1/3 m-4  border">
          <label htmlFor="name" className="text-white font-semibold">
            Name
          </label>
          <input
            className="border  rounded rounded:lg p-2 mb-2"
            type="text"
            placeholder="Enter Your name "
            onChange={(e) => setName(e.target.value)}
          />

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

          <button
            className=" text-white bg-green-800   flex items-center  rounded-lg   font-semibold border w-fit px-3 py-2 mt-2 mx-auto"
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
