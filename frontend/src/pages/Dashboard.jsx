import React, { useEffect, useState } from "react";
import Modal from "../components/Modal";
import CreateBtn from "../components/CreateBtn";
import { handleExportCSV } from "../services/ExportCsvService";
import DashboardStats from "../components/DashboardStats";
import ProductTable from "../components/ProductTable";
import { getProducts,createProduct, updateProduct ,deleteProduct} from "../services/productService";

function Dashboard() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [filter, setFilter] = useState("all"); // "all" or "low_stock"
  const [editingId, setEditingId] = useState(null);
  const [isModalOpen, setModalOpen] = useState(false);
  const [isDeleteOpen, setDeleteOpen] = useState(false);
  const [productToDelete, setProductToDelete] = useState(null);
  const [formData, setFormData] = useState({
    product_name: "",
    sku: "",
    quantity: 0,
    status: "in_stock",
    location: "",
  });

  // Function to open modal for adding new product
  const handleAddClick = () => {
    setFormData({
      product_name: "",
      sku: "",
      quantity: 0,
      status: "in_stock",
      location: "",
    });
    setModalOpen(true);
  };

  // function to handle form submission for both add and edit
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setLoading(true);
      if (editingId) {
        await updateProduct(editingId, formData);
      } else {
        await createProduct(formData);
      }
      setModalOpen(false);
      setEditingId(null);
      await fetchProducts();
    } catch (err) {
      console.error(err);
    }
  };

  // Function to fetch products from backend
  const fetchProducts = async () => {
    try {
      setLoading(true);
      const query = filter === "low_stock" ? "?status=low_stock" : "";
      const data = await getProducts(query);
      setProducts(data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };
  

//   Fetch products on component mount and when filter changes
  useEffect(() => {
    setLoading(true);
    fetchProducts();
  }, [filter]);

//  Calculate stats
  const totalStock = products.reduce((sum, p) => sum + p.quantity, 0);
  const lowStockCount = products.filter((p) => p.status === "low_stock").length;

  const handleEdit = (product) => {
    // Open the modal and populate form data
    setFormData({
      product_name: product.product_name,
      sku: product.sku,
      quantity: product.quantity,
      status: product.status,
      location: product.location,
    });
    setEditingId(product.id); // store which product is being edited
    setModalOpen(true);
  };

//   function to handle delete
  const handleDelete = (product) => {
    setProductToDelete(product);
    setDeleteOpen(true);
  };
  const handleConfirmDelete = async () => {
    try {
      await deleteProduct(productToDelete.id);
      setDeleteOpen(false);
      fetchProducts(); // refresh list
    } catch (err) {
      console.error(err);
    }
  };

// Export CSV handler
  const handleExport = () => {
    handleExportCSV();
  };

  if (loading) return <p className="text-center p-4">Loading...</p>;

  return (
    <div className="p-6">
      {/* Header + Add Button */}
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">📦 Product Dashboard</h1>
        <CreateBtn onClick={handleAddClick} />
        {/* Modal form for Creating product */}
        <Modal
          isOpen={isModalOpen}
          onClose={() => setModalOpen(false)}
          title="Add Product"
        >
          <form onSubmit={handleSubmit} className="flex flex-col gap-3">
            <input
              type="text"
              placeholder="Product Name"
              value={formData.product_name}
              onChange={(e) =>
                setFormData({ ...formData, product_name: e.target.value })
              }
              className="border px-2 py-1 rounded"
              required
            />
            <input
              type="text"
              placeholder="SKU"
              value={formData.sku}
              onChange={(e) =>
                setFormData({ ...formData, sku: e.target.value })
              }
              className="border px-2 py-1 rounded"
              required
            />
            <input
              type="number"
              placeholder="Quantity"
              value={formData.quantity}
              onChange={(e) =>
                setFormData({ ...formData, quantity: Number(e.target.value) })
              }
              className="border px-2 py-1 rounded"
              required
            />
            <input
              type="text"
              placeholder="Location"
              value={formData.location}
              onChange={(e) =>
                setFormData({ ...formData, location: e.target.value })
              }
              className="border px-2 py-1 rounded"
              required
            />
            <select
              value={formData.status}
              onChange={(e) =>
                setFormData({ ...formData, status: e.target.value })
              }
              className="border px-2 py-1 rounded"
            >
              <option value="in_stock">In Stock</option>
              <option value="low_stock">Low Stock</option>
              <option value="out_of_stock">Out of Stock</option>
            </select>
            <button
              type="submit"
              className="bg-green-500 text-white px-4 py-2 rounded"
            >
              Save
            </button>
          </form>
        </Modal>
      </div>

      {/* Dashboard Stats */}
      <DashboardStats
        products={products}
        totalStock={totalStock}
        lowStockCount={lowStockCount}
      />

      {/* Filter */}
      <div className="mb-4">
        <button
          className={`px-3 py-1 rounded mr-2 ${
            filter === "all" ? "bg-blue-600 text-white" : "bg-gray-200"
          }`}
          onClick={() => setFilter("all")}
        >
          All
        </button>
        <button
          className={`px-3 py-1 rounded ${
            filter === "low_stock" ? "bg-blue-600 text-white" : "bg-gray-200"
          }`}
          onClick={() => setFilter("low_stock")}
        >
          Low Stock
        </button>
      </div>

      {/* Product Table */}
      <ProductTable
        products={products}
        handleEdit={handleEdit}
        handleDelete={handleDelete}
      />

      <Modal isOpen={isDeleteOpen} onClose={() => setDeleteOpen(false)} title="Confirm Delete">
  <p>Are you sure you want to delete <strong>{productToDelete?.product_name}</strong>?</p>
  <div className="flex justify-end gap-2 mt-4">
    <button
      className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
      onClick={() => setDeleteOpen(false)}
    >
      Cancel
    </button>
    <button
      className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
      onClick={handleConfirmDelete}
    >
      Delete
    </button>
  </div>
</Modal>

      {/* Export Button */}
      <div className="flex justify-end mt-4">

      <button
        className="bg-yellow-500 hover:bg-yellow-600 text-white px-2 py-1  rounded"
        onClick={() => handleExport()}
        >
        Download CSV
      </button>
        </div>
    </div>
  );
}

export default Dashboard;
