import axios from "axios";
import { Button, TextInput } from "flowbite-react";
import React, { useState } from "react";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const baseURL = "http://localhost:8000/v1/exchanges";
const Header = () => {
  const [values, setValues] = useState({
    api_key: "",
    secret_key: "",
  });
  const data = [
    {
      id: 1,
      name: "api_key",
      type: "password",
      placeholder: "APIKEY",
      // required: true,
    },
    {
      id: 2,
      name: "secret_key",
      type: "password",
      placeholder: "SECRETKEY",
      // required: true,
    },
  ];

  const keys = {
    api_key: values.api_key,
    secret_key: values.secret_key,
  };

  const config = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  const onChange = (e) => {
    setValues({ ...values, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    axios.put(baseURL, keys, config).then((res) => {
      console.log(res.data);
    });
  };
  return (
    <div className="flex mt-3 mb-20 justify-between mx-5 items-center ">
      <h1 className="text-3xl font-bold underline text-[#F3c535]">
        NOT HUMAN TRADE
      </h1>
      {data.map((i) => (
        <>
          <span className="mx-[5rem] text-[#848e9c]">{i.name}</span>
          <TextInput
            {...i}
            key={i.id}
            onChange={onChange}
            className="appearance-none bg-transparent border-none w-96 text-amber-50 text-lg font-semibold mr-3 py-1 px-2 leading-tight focus:outline-none"
          />
        </>
      ))}

      <Button
        gradientMonochrome="lime"
        type="submit"
        onClick={(e) => handleSubmit(e)}
        className="px-6 bg-[#F3c535] hover:bg-[#F3c535] "
      >
        <span className="text-[#181a20]">SAVE</span>
      </Button>
      <ToastContainer
        position="bottom-right"
        autoClose={2500}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="dark"
      />
    </div>
  );
};

export default Header;
