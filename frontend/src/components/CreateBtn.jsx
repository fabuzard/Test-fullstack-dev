import React from 'react'

function CreateBtn({ onClick }) {
  return (
    <button
      onClick={onClick}
      className="bg-blue-600 hover:bg-blue-700 text-white font-semibold px-4 py-2 rounded shadow-md transition duration-200"
    >
      + Add New
    </button>
  )
}

export default CreateBtn
