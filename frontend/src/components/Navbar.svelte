<nav class="navbar">
    <div class="logo">Go file manager</div>
    <ul>
        <li><a href="/" class="logout">Log out</a></li>
    </ul>
</nav>

<style>
    .navbar{
        display: flex;
        width: 100%;
        height: 6rem;
        background-color: var(--gray);
        justify-content: space-between;
        align-items: center;
        color: var(--text-primary);
        box-shadow: rgba(0, 0, 0, 0.25) 0px 54px 55px, rgba(0, 0, 0, 0.12) 0px -12px 30px, rgba(0, 0, 0, 0.12) 0px 4px 6px, rgba(0, 0, 0, 0.17) 0px 12px 13px, rgba(0, 0, 0, 0.09) 0px -3px 5px;
    }
    .logo{
        margin-left: 10rem;
        font-size: 2rem;
        cursor: pointer;
        color: var(--blue);
    }
    ul{
        list-style: none;
    }
    a{
        text-decoration: none;
    }
    .logout{
        margin-right: 5rem;
        font-size: 1.25rem;
    }
</style>