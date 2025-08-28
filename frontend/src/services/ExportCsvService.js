const API_URL = "https://inventory-backend-fabuzard2-b3f3496681aa.herokuapp.com:8080"; 

export const handleExportCSV = async () => {
  const res = await fetch(`${API_URL}/products/export`);
  const blob = await res.blob();
  const url = window.URL.createObjectURL(blob);

  const a = document.createElement("a");
  a.href = url;
  a.download = "products.csv";
  a.click();
  window.URL.revokeObjectURL(url);
};