import axios from "axios";

export async function LoginService(form, navigate) {
    try {
        const res = await axios.post("http://localhost:8080/login", form);
        localStorage.setItem("token", res.data.token);
        navigate("/dashboard");
    } catch (err) {
        console.error(err);
        alert(err.response?.data?.error || "Login failed");
    }
}

export async function RegisterService(form, navigate) {
    try {
        await axios.post("http://localhost:8080/register", form);
        alert("Registered successfully! Please login.");
        navigate("/login");
      } catch (err) {
        console.error(err);
        alert(err.response?.data?.error || "Registration failed");
      }
    }
