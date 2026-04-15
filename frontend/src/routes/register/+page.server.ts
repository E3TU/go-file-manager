import { redirect, type Actions } from '@sveltejs/kit';

const API_BASE = 'http://localhost:8080/api';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const name = formData.get('name')?.toString() || '';
		const email = formData.get('email')?.toString() || '';
		const password = formData.get('password')?.toString() || '';

		const res = await fetch(`${API_BASE}/auth/register`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ name, email, password })
		});

		const data = await res.json();

		if (!res.ok) {
			return { error: data.error || 'Registration failed' };
		}

		// Set cookie from the response session (proxying like login)
		cookies.set('a_session', data.session.sessionSecret, {
			path: '/',
			httpOnly: true,
			sameSite: 'lax',
			secure: false, // true in prod
			maxAge: 60 * 60, // 1 hour
		});

		throw redirect(303, '/');
	}
};