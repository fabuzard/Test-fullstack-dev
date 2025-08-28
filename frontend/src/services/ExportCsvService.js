export const handleExportCSV = async () => {
    const res = await fetch("http://localhost:8080/products/export");
    const blob = await res.blob();
    const url = window.URL.createObjectURL(blob);
  
    const a = document.createElement("a");
    a.href = url;
    a.download = "products.csv";
    a.click();
    window.URL.revokeObjectURL(url);
  };
  