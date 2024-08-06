document.addEventListener('DOMContentLoaded', function() {
    const sidebar = document.getElementById('sidebar');
    const content = document.getElementById('content');
    const sidebarToggle = document.getElementById('sidebarToggle');

    // Load sidebar state from localStorage
    const sidebarState = localStorage.getItem('sidebarState');
    if (sidebarState === 'active') {
        sidebar.classList.add('active');
        content.classList.add('active');
    }

    sidebarToggle.addEventListener('click', function() {
        sidebar.classList.toggle('active');
        content.classList.toggle('active');

        // Save sidebar state to localStorage
        if (sidebar.classList.contains('active')) {
            localStorage.setItem('sidebarState', 'active');
        } else {
            localStorage.removeItem('sidebarState');
        }
    });
});