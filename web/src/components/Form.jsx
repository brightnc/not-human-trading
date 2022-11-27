import { TextInput } from "flowbite-react";
import React, { useState } from "react";

const Form = (props) => {
  const { errorMessage, onChange, id, ...inputProps } = props;

  return (
    <form className="w-full max-w-[20%]">
      <div className="flex flex-col items-center border-b border-teal-500 py-2">
        <TextInput
          {...inputProps}
          required={true}
          onChange={onChange}
          className="appearance-none bg-transparent border-none w-full text-amber-50 text-lg font-semibold mr-3 py-1 px-2 leading-tight focus:outline-none"
        />
      </div>
    </form>
  );
};

export default Form;
