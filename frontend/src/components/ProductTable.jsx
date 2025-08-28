import React from 'react'
import Barcode from './Barcode'
function ProductTable({ products, handleEdit, handleDelete }) {
  return (
    <div className="overflow-x-auto">
    <table className="min-w-full border border-gray-200 shadow-lg rounded-lg">
      <thead className="bg-gray-100 text-gray-600 uppercase text-sm">
        <tr>
          <th className="px-4 py-2 border">ID</th>
          <th className="px-4 py-2 border">Name</th>
          <th className="px-4 py-2 border">SKU</th>
          <th className="px-4 py-2 border">Quantity</th>
          <th className="px-4 py-2 border">Status</th>
          <th className="px-4 py-2 border">Location</th>
          <th className="px-4 py-2 border">Actions</th>
          <th className="px-4 py-2 border">Barcode</th>
        </tr>
      </thead>
      <tbody>
        {products.map((p) => (
          <tr key={p.id} className="hover:bg-gray-50">
            <td className="px-4 py-2 border">{p.id}</td>
            <td className="px-4 py-2 border">{p.product_name}</td>
            <td className="px-4 py-2 border">{p.sku}</td>
            <td className="px-4 py-2 border">{p.quantity}</td>
            <td className="px-4 py-2 border">
              <span
                className={`px-2 py-1 rounded text-white text-xs ${
                  p.status === "in_stock"
                    ? "bg-green-500"
                    : p.status === "low_stock"
                    ? "bg-yellow-500"
                    : "bg-red-500"
                }`}
              >
                {p.status.replace("_", " ")}
              </span>
            </td>
            <td className="px-4 py-2 border">{p.location}</td>
            <td className="px-4 py-2 border">
              <div className='flex gap-2 justify-center'>

              <button
                className="bg-yellow-500 hover:bg-yellow-600 text-white px-2 py-1 rounded"
                onClick={() => handleEdit(p)}
                >
                Edit
              </button>
              <button
                className="bg-red-500 hover:bg-red-600 text-white px-2 py-1 rounded"
                onClick={() => handleDelete(p)}
                >
                Delete
              </button>
                </div>
            </td>
            <td className='px-4 py-2 border flex justify-center'><Barcode sku={p.sku}/></td>
          </tr>
        ))}
      </tbody>
    </table>
  
  </div>
  )
}

export default ProductTable
