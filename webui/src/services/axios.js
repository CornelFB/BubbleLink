import axios from "axios";
import { useAuth } from "@/composables/useAuth"; // import your existing composable

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

// automatically attach Bearer token if logged in
instance.interceptors.request.use((config) => {
	try {
		const { user } = useAuth();
		const token = user.value?.apiKey;
		if (token) config.headers.Authorization = `Bearer ${token}`;
	} catch {
		// ignore if composable not ready
	}
	return config;
});

export default instance;
