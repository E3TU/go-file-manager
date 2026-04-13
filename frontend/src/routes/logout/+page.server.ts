import { redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ cookies }) => {
		cookies.delete('a_session', { path: '/' });
		redirect(303, '/login');
	}
};