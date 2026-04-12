<script lang="ts">
    import Navbar from '../../components/Navbar.svelte';
    import { login, setSession } from '$lib/api';

    let email = $state('');
    let password = $state('');
    let error = $state('');
    let loading = $state(false);

    async function handleSubmit(e: Event) {
        e.preventDefault();
        error = '';
        loading = true;
        try {
            const session = await login(email, password);
            setSession(session.secret);
            window.location.href = '/';
        } catch (err) {
            error = err instanceof Error ? err.message : 'Login failed';
        } finally {
            loading = false;
        }
    }
</script>

<Navbar />
<div class="container">
    <div class="form-card">
        <h2>Log In</h2>
        <form onsubmit={handleSubmit}>
            <div class="input-group">
                <input type="email" placeholder="Email" bind:value={email} required />
            </div>
            <div class="input-group">
                <input type="password" placeholder="Password" bind:value={password} required />
            </div>
            {#if error}
                <p class="error">{error}</p>
            {/if}
            <button type="submit" class="submit-btn" disabled={loading}>
                {loading ? 'Logging in...' : 'Log In'}
            </button>
        </form>
        <p class="switch">Don't have an account? <a href="/register">Register</a></p>
    </div>
</div>

<style>
    .container {
        display: flex;
        align-items: center;
        justify-content: center;
        height: calc(100vh - 6rem);
        width: 100%;
    }
    .form-card {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 25rem;
        background-color: var(--gray);
        box-shadow:
            rgba(0, 0, 0, 0.16) 0px 10px 36px 0px,
            rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
        border-radius: 16px;
        padding: 2rem;
    }
    h2 {
        color: var(--text-primary);
        font-size: 1.75rem;
        margin-bottom: 1.5rem;
    }
    form {
        display: flex;
        flex-direction: column;
        width: 100%;
        gap: 1rem;
    }
    .input-group input {
        width: 100%;
        padding: 0.75rem 1rem;
        border: 2px solid #333;
        border-radius: 8px;
        background-color: #191919;
        color: var(--text-primary);
        font-size: 1rem;
        font-family: inherit;
        box-sizing: border-box;
    }
    .input-group input:focus {
        outline: none;
        border-color: var(--blue);
    }
    .input-group input::placeholder {
        color: #666;
    }
    .submit-btn {
        margin-top: 0.5rem;
        padding: 0.75rem;
        font-size: 1rem;
        font-family: inherit;
        border: 2px solid var(--blue);
        border-radius: 8px;
        background-color: var(--blue);
        color: var(--text-primary);
        transition: 0.3s;
        cursor: pointer;
    }
    .submit-btn:hover {
        background-color: transparent;
        transition: 0.3s;
    }
    .switch {
        margin-top: 1.5rem;
        color: var(--text-primary);
        font-size: 0.9rem;
    }
    .switch a {
        color: var(--blue);
        text-decoration: none;
    }
    .switch a:hover {
        text-decoration: underline;
    }
    .error {
        color: #ff4444;
        font-size: 0.9rem;
        text-align: center;
    }
</style>