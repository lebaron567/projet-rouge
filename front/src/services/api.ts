import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api', // Remplacez par l'URL de votre API Go
});

export const fetchTweets = async () => {
  const response = await api.get('/tweets');
  return response.data;
};

export default api;