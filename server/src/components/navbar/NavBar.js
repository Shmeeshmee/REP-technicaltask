import "./NavBar.css";
function NavBar() {
  return (
    <div className="navbar">
        <a href="/add_pack">ADD PACK</a>
        <a href="/delete_pack">DELETE PACK</a>
        <a href="/get_all_packs">GET ALL PACKS</a>
      <a href="/order_packs">ORDER PACKS</a>
    </div>
  );
}
export default NavBar;
