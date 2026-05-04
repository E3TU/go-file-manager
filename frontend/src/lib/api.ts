const API_BASE: string = 'http://localhost:8080/api';

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
        body: JSON.stringify({ name, email, password }),
        credentials: 'include'
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

export interface UploadedFile {
    id: string;
    name: string;
    sizeOriginal: number;
    mimeType: string;
    createdAt: string;
}

export async function uploadFile(file: File): Promise<UploadedFile> {
    const formData = new FormData();
    formData.append('file', file);

    const res = await fetch(`${API_BASE}/storage/files`, {
        method: 'POST',
        body: formData,
        credentials: 'include'
    });
    if (!res.ok) {
        const err = await res.json();
        throw new Error(err.error || 'Upload failed');
    }
    return res.json();
}

export async function listFiles(): Promise<{ files: UploadedFile[] }> {
    const res = await fetch(`${API_BASE}/storage/files`, {
        credentials: 'include'
    });
    if (!res.ok) {
        throw new Error('Failed to list files');
    }
    return res.json();
}

export async function deleteFile(fileId: string): Promise<void> {
    const res = await fetch(`${API_BASE}/storage/files/${fileId}`, {
        method: 'DELETE',
        credentials: 'include'
    });
    if (!res.ok) {
        throw new Error('Delete failed');
    }
}

export async function downloadFile(fileId: string, fileName: string): Promise<void> {
    const res = await fetch(`${API_BASE}/storage/files/${fileId}/download`, {
        credentials: 'include'
    });
    if (!res.ok) {
        throw new Error('Download failed');
    }
    const blob = await res.blob();
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = fileName;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
}