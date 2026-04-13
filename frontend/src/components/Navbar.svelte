<script>
	import { page } from '$app/stores';
	import { enhance } from '$app/forms';
</script>

<nav class="navbar">
	<div class="logo">Go file manager</div>
	<ul>
			<li>
				<form method="POST" action="/logout" use:enhance>
					<button type="submit" class="logout">Log out</button>
				</form>
			</li>
	</ul>
</nav>

<style>
	.navbar {
		display: flex;
		width: 100%;
		height: 6rem;
		background-color: var(--gray);
		justify-content: space-between;
		align-items: center;
		color: var(--text-primary);
		box-shadow:
			rgba(0, 0, 0, 0.25) 0px 54px 55px,
			rgba(0, 0, 0, 0.12) 0px -12px 30px,
			rgba(0, 0, 0, 0.12) 0px 4px 6px,
			rgba(0, 0, 0, 0.17) 0px 12px 13px,
			rgba(0, 0, 0, 0.09) 0px -3px 5px;
	}
	.logo {
		margin-left: 10rem;
		font-size: 2rem;
		cursor: pointer;
		color: var(--blue);
	}
	ul {
		list-style: none;
	}
	.logout {
		margin-right: 5rem;
		font-size: 1.25rem;
		background: none;
		border: none;
		color: var(--text-primary);
		cursor: pointer;
		font-family: inherit;
		transition: 0.5;
	}
	.logout:hover {
		color: var(--blue);
		transition: 0.5;
	}
</style>
