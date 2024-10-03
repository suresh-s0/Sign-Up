import { useNavigate } from 'react-router-dom';

function SignIn(){

    return(

        <div className="flex flex-col  justify-center items-center  sm:w-full md:w-2/3 lg:w-1/2 xl:w-full border-2  px-6 py-24 ">
        <h1 className="font-bold text-4xl text-white">Sign-In </h1>

        <div className="flex flex-col justify-center  p-3  sm:w-full md:w-1/3 lg:w-1/2 xl:w-1/3 m-4  border">
          
          <label htmlFor="email" className="text-white font-semibold">
            Email
          </label>
          <input
            className="border  rounded rounded:md p-2 mb-2"
            type="email"
            placeholder="Enter Your Email"
            // onChange={(e) => setEmail(e.target.value)}
          />

          <label htmlFor="pasword" className="text-white font-semibold">
            Password
          </label>
          <input
            className="border  rounded-md p-2 mb-2"
            type="password"
            placeholder="Enter Your password"
            // onChange={(e) => setPass(e.target.value)}
          />

          <button
            className=" text-white bg-green-800   flex items-center  rounded-lg   font-semibold border w-fit px-4 py-2 mx-auto"
            type="button"
            // onClick={handleclic}
          >
            Submit
          </button>

         
        </div>
      </div>
    );

}

export default SignIn;