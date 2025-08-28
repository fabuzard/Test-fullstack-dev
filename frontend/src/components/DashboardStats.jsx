import React from 'react'

function DashboardStats({ products, totalStock, lowStockCount }) {
  return (
    <div className="flex gap-4 mb-4">
    <div className="bg-white p-4 rounded shadow text-center flex-1">
      <p className="text-gray-500">Total Products</p>
      <p className="text-xl font-bold">{products.length}</p>
    </div>
    <div className="bg-white p-4 rounded shadow text-center flex-1">
      <p className="text-gray-500">Total Stock</p>
      <p className="text-xl font-bold">{totalStock}</p>
    </div>
    <div className="bg-white p-4 rounded shadow text-center flex-1">
      <p className="text-gray-500">Low Stock</p>
      <p className="text-xl font-bold">{lowStockCount}</p>
    </div>
  </div>
  )
}

export default DashboardStats
