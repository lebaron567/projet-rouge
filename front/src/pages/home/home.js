import Navbar from "../../components/Navbar/Navbar";
import './home.css';

const Home = () => {
    return (
        <>
        <Navbar/>
        <div className="home">
            <h1>Home</h1>
            <p>Welcome to the home page!</p>
        </div>
        </>
    );
}
export default Home;