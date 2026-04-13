import { redirect, type Handle } from '@sveltejs/kit';

const API_BASE = 'http://localhost:8080/api';
const PUBLIC_ROUTES = ['/login', '/register', '/logout'];
const COOKIE_NAME = 'a_session';

export const handle: Handle = async ({ event, resolve }) => {
	const pathname = event.url.pathname;

	const isPublicRoute = PUBLIC_ROUTES.some(
		(route) => pathname === route || pathname.startsWith(route + '/')
	);

	if (isPublicRoute) {
		return resolve(event);
	}

	const sessionCookie = event.cookies.get(COOKIE_NAME);

	if (!sessionCookie) {
		redirect(303, '/login');
	}

	try {
		const res = await fetch(`${API_BASE}/auth/session`, {
			headers: {
				Cookie: `${COOKIE_NAME}=${sessionCookie}`
			}
		});

		if (!res.ok) {
			redirect(303, '/login');
		}

		const session = await res.json();

		if (!session.valid) {
			redirect(303, '/login');
		}

		event.locals.user = {
			userId: session.userId,
			name: session.name,
			email: session.email
		};
	} catch {
		redirect(303, '/login');
	}

	return resolve(event);
};