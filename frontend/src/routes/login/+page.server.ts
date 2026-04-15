import { redirect, type Actions } from '@sveltejs/kit';

const API_BASE = 'http://localhost:8080/api';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const email = formData.get('email')?.toString() || '';
		const password = formData.get('password')?.toString() || '';

		const res = await fetch(`${API_BASE}/auth/session`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email, password }),
		});

		const data = await res.json();

		if (!res.ok) {
			return { error: data.error || 'Login failed' };
		}

		// Store the session secret in a cookie from SvelteKit
		cookies.set('a_session', data.sessionSecret, {
			path: '/',
			httpOnly: true,
			sameSite: 'lax',
			secure: false, // true in production
			maxAge: 60 * 60, // 1 hour
		});

		throw redirect(303, '/');
	}
};