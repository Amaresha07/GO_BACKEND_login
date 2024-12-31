<script lang="ts">
    let email = "";
    let password = "";
    let confirmPassword = "";
    let isLoading = false;
    let isSignup = false;

    const handleSubmit = async () => {
        if (isSignup && password !== confirmPassword) {
            alert("Passwords do not match");
            return;
        }

        isLoading = true;

        try {
            const endpoint = isSignup
                ? "http://localhost:8000/register"
                : "http://localhost:8000/login";

            const response = await fetch(endpoint, {
                method: "POST",
                body: JSON.stringify({
                    email: email,
                    password: password,
                }),
            });

            if (response.ok) {
                const data = await response.json();
                alert(
                    data.message ||
                        `${isSignup ? "Signup" : "Login"} successful`,
                );
            } else {
                const errorData = await response.json();
                alert(
                    errorData.message ||
                        `${isSignup ? "Signup" : "Login"} failed`,
                );
            }
        } catch (error) {
            console.error(error); // Log error for debugging
            alert("An error occurred while connecting to the server");
        } finally {
            isLoading = false;
        }
    };

    const toggleSignup = () => {
        isSignup = !isSignup;
    };
</script>

<div class="login-container">
    <div class="login-left">
        <div class="brand-showcase">
            <h1>Welcome to Our Platform</h1>
            <p>Experience the next generation of secure login</p>
        </div>
    </div>

    <div class="login-right">
        <div class="login-card">
            <div class="brand">
                <h2>{isSignup ? "Create an Account" : "Welcome Back"}</h2>
                <p>
                    {isSignup
                        ? "Sign up to get started"
                        : "Please sign in to continue"}
                </p>
            </div>

            <form on:submit|preventDefault={handleSubmit} class="login-form">
                <div class="input-group">
                    <label for="email">Email</label>
                    <input
                        type="email"
                        id="email"
                        bind:value={email}
                        placeholder="your@gmail.com"
                        required
                    />
                </div>

                <div class="input-group">
                    <label for="password">Password</label>
                    <input
                        type="password"
                        id="password"
                        bind:value={password}
                        placeholder="password"
                        required
                    />
                </div>

                {#if isSignup}
                    <div class="input-group">
                        <label for="confirmPassword">Confirm Password</label>
                        <input
                            type="password"
                            id="confirmPassword"
                            bind:value={confirmPassword}
                            placeholder="confirmPassword"
                            required
                        />
                    </div>
                {/if}

                <button type="submit" class="login-button" disabled={isLoading}>
                    {#if isLoading}
                        <span class="loader"></span>
                    {:else}
                        {isSignup ? "Sign Up" : "Sign In"}
                    {/if}
                </button>
            </form>

            <div class="signup-link">
                {isSignup
                    ? "Already have an account?"
                    : "Don't have an account?"}
                <a href="#" on:click|preventDefault={toggleSignup}>
                    {isSignup ? "Log in" : "Sign up"}
                </a>
            </div>
        </div>
    </div>
</div>

<style>
    body {
        font-family: "Roboto", sans-serif;
        margin: 0;
        padding: 0;
        background: linear-gradient(135deg, #ff7e5f, #feb47b);
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .login-container {
        display: flex;
        max-width: 900px;
        width: 100%;
        background: #ffffff;
        border-radius: 10px;
        overflow: hidden;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.2);
    }

    .login-left {
        flex: 1;
        background: linear-gradient(135deg, #0072ff, #00c6ff);
        color: white;
        padding: 40px;
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    .brand-showcase h1 {
        margin: 0;
        font-size: 2.5em;
    }

    .brand-showcase p {
        margin-top: 10px;
        font-size: 1.2em;
    }

    .login-right {
        flex: 1;
        padding: 40px;
    }

    .login-card {
        max-width: 400px;
        margin: 0 auto;
    }

    .brand h2 {
        margin: 0;
        font-size: 2em;
        color: #333;
    }

    .brand p {
        margin-top: 10px;
        color: #666;
    }

    .login-form {
        margin-top: 20px;
    }

    .input-group {
        margin-bottom: 20px;
    }

    .input-group label {
        display: block;
        font-weight: bold;
        margin-bottom: 5px;
        color: #444;
    }

    .input-group input {
        width: 100%;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 5px;
        font-size: 1em;
    }

    .login-button {
        width: 100%;
        padding: 12px;
        background: linear-gradient(135deg, #6a11cb, #2575fc);
        color: white;
        font-weight: bold;
        font-size: 1em;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        transition: background 0.3s ease;
    }

    .login-button:hover {
        background: linear-gradient(135deg, #2575fc, #6a11cb);
    }

    .login-button:disabled {
        background: #ccc;
        cursor: not-allowed;
    }

    .signup-link {
        margin-top: 20px;
        text-align: center;
        color: #444;
    }

    .signup-link a {
        color: #007bff;
        text-decoration: none;
        font-weight: bold;
    }

    .signup-link a:hover {
        text-decoration: underline;
    }

    .loader {
        border: 3px solid #f3f3f3;
        border-radius: 50%;
        border-top: 3px solid #3498db;
        width: 14px;
        height: 14px;
        animation: spin 2s linear infinite;
        display: inline-block;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
</style>
