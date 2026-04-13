const API_BASE : string = 'http://localhost:8080/api';

export async function login(email: string, password: string) {
    const res = await fetch(`${API_BASE}/auth/session`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
        credentials: 'include'
    });
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || 'Login failed');
    }
    return res.json();
}

export async function register(name: string, email: string, password: string) {
    const res = await fetch(`${API_BASE}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password })
    });
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || 'Registration failed');
    }
    return res.json();
}

export function getSession(): string | null {
    return localStorage.getItem('session');
}

export function setSession(session: string) {
    localStorage.setItem('session', session);
}

export function clearSession() {
    localStorage.removeItem('session');
}