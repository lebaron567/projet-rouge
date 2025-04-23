import './navbar.css';

const Navbar = () => {
    return (
        <nav className="navbar">
        <h1 className="logo">Y</h1>
        <div className="links">
            <a href="/">Home</a>
            <a href="/explore">Explore</a>
            <a href="/notifications">Notifications</a>
            <a href="/bookmarks">Bookmarks</a>
            <a href="/profile">Profile</a>
            <a href="/more">More</a>
            <button href="/create">Post</button>
        </div>
        </nav>
    );
}

export default Navbar;