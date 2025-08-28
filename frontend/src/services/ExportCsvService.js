const API_URL = "https://inventory-backend-fabuzard2-b3f3496681aa.herokuapp.com"; 

function getAuthHeaders() {
  const token = localStorage.getItem("token");
  return {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
}

export const handleExportCSV = async () => {
  const res = await fetch(`${API_URL}/products/export`,getAuthHeaders());
  const blob = await res.blob();
  const url = window.URL.createObjectURL(blob);

  const a = document.createElement("a");
  a.href = url;
  a.download = "products.csv";
  a.click();
  window.URL.revokeObjectURL(url);
};