<header class="navbar">
    <div class="nav-container">
        <a href="/" class="brand-logo">Travel<span>Sphere</span></a>
        <nav class="nav-links">
            <a href="/" class="active">Home</a>
            <a href="/countries">Countries</a>
            <a href="/wishlist">Wishlist</a>
            <a href="/dashboard">Dashboard</a>
        </nav>
        
        <div class="auth-actions">
            {{if .IsLoggedIn}}
                <div class="user-profile-status" style="display: flex; align-items: center; gap: 12px; font-family: system-ui, sans-serif;">
                    <span class="welcome-msg" style="color: #1e293b; font-weight: 600; font-size: 14px;">Welcome {{.CurrentUserName}}!</span>
                    <a href="/logout" class="logout-btn" style="text-decoration: none; background-color: #f1f5f9; color: #64748b; padding: 6px 12px; border-radius: 6px; font-size: 13px; font-weight: 500; border: 1px solid #e2e8f0;">Logout</a>
                </div>
            {{else}}
                <button type="button" class="login-btn" onclick="openLoginModal()" style="cursor: pointer; font-size: 13px; border: none; padding: 6px 16px; background-color: #4f46e5; color: white; border-radius: 6px; font-weight: 500;">
                    Login
                </button>
            {{end}}
        </div>
    </div>
</header>

<div id="loginModal" style="display: none; position: fixed; z-index: 1000; left: 0; top: 0; width: 100%; height: 100%; background-color: rgba(15, 23, 42, 0.6); backdrop-filter: blur(4px); align-items: center; justify-content: center;">
    <div style="background-color: #ffffff; padding: 32px; border-radius: 16px; border: 1px solid #e2e8f0; max-width: 400px; width: 90%; box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1); position: relative; font-family: system-ui, sans-serif;">
        <span onclick="closeLoginModal()" style="position: absolute; right: 20px; top: 16px; font-size: 20px; color: #94a3b8; cursor: pointer; font-weight: bold;">&times;</span>
        
        <h2 style="margin: 0 0 8px 0; font-size: 24px; font-weight: 700; color: #0f172a;">Sign In</h2>
       
        <form action="/login" method="POST" style="display: flex; flex-direction: column; gap: 16px; margin: 0;">
            <div>
                <label style="display: block; font-size: 12px; font-weight: 600; color: #475569; margin-bottom: 6px;">TRAVELER NAME</label>
                <input type="text" name="username" placeholder="e.g. Jerry" required 
                       style="width: 100%; padding: 10px 12px; border: 1px solid #cbd5e1; border-radius: 8px; font-size: 14px; outline: none; box-sizing: border-box; background-color: #f8fafc;">
            </div>
            <button type="submit" style="width: 100%; background-color: #4f46e5; color: #ffffff; border: none; padding: 12px; font-weight: 600; border-radius: 8px; cursor: pointer; font-size: 14px; margin-top: 4px;">
            LOGIN
            </button>
        </form>
    </div>
</div>

<script>
function openLoginModal() {
    document.getElementById("loginModal").style.display = "flex";
}
function closeLoginModal() {
    document.getElementById("loginModal").style.display = "none";
}
// Automatically hide the screen overlay view grid if click triggers outside the dialog boundaries
window.onclick = function(event) {
    const modal = document.getElementById("loginModal");
    if (event.target == modal) {
        closeLoginModal();
    }
}
</script>