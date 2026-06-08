<div style="padding: 60px 24px; font-family: system-ui, -apple-system, sans-serif; text-align: center; background-color: #f8fafc; min-height: 100vh; display: flex; flex-direction: column; align-items: center; justify-content: center;">
    <div style="background: #ffffff; padding: 48px; border-radius: 20px; border: 1px solid #e2e8f0; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); max-width: 480px; width: 100%;">
        <span style="font-size: 36px; display: block; color:red; margin-bottom: 16px;">404 error</span>
        <h1 style="font-size: 36px; font-weight: 800; color: #0f172a; margin: 0 0 12px 0;">Destination Not Found</h1>
        
        <p style="color: #64748b; font-size: 16px; line-height: 1.6; margin: 0 0 32px 0;">
            {{if .ErrorMessage}}
                {{.ErrorMessage}}
            {{else}}
                The travel destination or page you are looking for does not exist or has been moved.
            {{end}}
        </p>
        
        <a href="/countries" style="text-decoration: none; background-color: #4f46e5; color: #ffffff; padding: 12px 32px; font-weight: 600; border-radius: 8px; display: inline-block; font-size: 15px; transition: background 0.15s ease;">
            Back to Country Explorer
        </a>
    </div>
</div>