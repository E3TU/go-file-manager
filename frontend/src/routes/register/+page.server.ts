import { redirect, type Actions } from '@sveltejs/kit';

const API_BASE = 'http://localhost:8080/api';

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const formData = await request.formData();
		const name = formData.get('name')?.toString() || '';
		const email = formData.get('email')?.toString() || '';
		const password = formData.get('password')?.toString() || '';
		const confirmPassword = formData.get('confirmPassword')?.toString() || '';

		if (password !== confirmPassword) {
			return { error: 'Passwords do not match' };
		}

		const registerRes = await fetch(`${API_BASE}/auth/register`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ name, email, password })
		});

		if (!registerRes.ok) {
			const err = await registerRes.json();
			return { error: err.error || 'Registration failed' };
		}

		redirect(303, '/');
	}
};