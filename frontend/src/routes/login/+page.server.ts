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
			body: JSON.stringify({ email, password })
		});

		if (!res.ok) {
			const err = await res.json();
			return { error: err.error || 'Login failed' };
		}

		redirect(303, '/');
	}
};