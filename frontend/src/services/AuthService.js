import axios from "axios";

const API_URL = "https://inventory-backend-fabuzard2-b3f3496681aa.herokuapp.com"; 

export async function LoginService(form, navigate) {
    try {
        const res = await axios.post(`${API_URL}/login`, form);
        localStorage.setItem("token", res.data.token);
        navigate("/dashboard");
    } catch (err) {
        console.error(err);
        alert(err.response?.data?.error || "Login failed");
    }
}

export async function RegisterService(form, navigate) {
    try {
        await axios.post(`${API_URL}/register`, form);
        alert("Registered successfully! Please login.");
        navigate("/login");
    } catch (err) {
        console.error(err);
        alert(err.response?.data?.error || "Registration failed");
    }
}