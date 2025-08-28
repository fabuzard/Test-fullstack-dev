import axios from "axios";

const API_URL = "http://34.101.244.42:8080/products";

// Helper to get headers with token
function getAuthHeaders() {
  const token = localStorage.getItem("token");
  return {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
}

export async function getProducts(filter = "") {
  try {
    const res = await axios.get(`${API_URL}${filter}`, getAuthHeaders());
    return res.data || [];
  } catch (err) {
    console.error("Error fetching products:", err);
    throw err;
  }
}

export async function createProduct(product) {
  try {
    const res = await axios.post(API_URL, product, getAuthHeaders());
    return res.data;
  } catch (err) {
    console.error("Error creating product:", err);
    throw err;
  }
}

export async function updateProduct(id, product) {
  try {
    const res = await axios.put(`${API_URL}/${id}`, product, getAuthHeaders());
    return res.data;
  } catch (err) {
    console.error("Error updating product:", err);
    throw err;
  }
}

export async function deleteProduct(id) {
  try {
    const res = await axios.delete(`${API_URL}/${id}`, getAuthHeaders());
    return res.data;
  } catch (err) {
    console.log(id);
    console.error("Error deleting product:", err);
    throw err;
  }
}
