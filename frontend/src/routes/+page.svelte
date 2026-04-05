<div class="container">
    <h1>Go file manager</h1>
    <div class="file-upload">
        <p>Upload a file</p>
    </div>
</div>

<style>
    .container{
        display: flex;
        align-items: center;
        flex-direction: column;
        height: 100vh;
        width: 100%;
    }
    h1{
        color: var(--text-primary);
        margin-top: 5rem;
        font-size: 3rem;
    }
    .file-upload{
        height: 20rem;
        width: 20rem;
        background-color: var(--gray);
        box-shadow: rgba(0, 0, 0, 0.16) 0px 10px 36px 0px, rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
        border-radius: 16px;
        margin-top: 2rem;
    }
    p{
        color: var(--text-primary);
    }
</style>