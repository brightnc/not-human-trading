import React from 'react'

const FormInput = (props) => {
  const {errorMessage, onChange, id, ...inputProps} = props
  return (
    <div>
        <input
            {...inputProps}
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <span>{errorMessage}</span>
    </div>
  )
}

export default FormInput