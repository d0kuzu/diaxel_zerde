export const validateEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

export const validatePassword = (password: string): { isValid: boolean; message?: string } => {
  if (password.length < 8) {
    return { isValid: false, message: 'Пароль должен содержать минимум 8 символов' };
  }
  
  if (!/(?=.*[a-z])/.test(password)) {
    return { isValid: false, message: 'Пароль должен содержать хотя бы одну строчную букву' };
  }
  
  if (!/(?=.*[A-Z])/.test(password)) {
    return { isValid: false, message: 'Пароль должен содержать хотя бы одну заглавную букву' };
  }
  
  if (!/(?=.*\d)/.test(password)) {
    return { isValid: false, message: 'Пароль должен содержать хотя бы одну цифру' };
  }
  
  return { isValid: true };
};

export const validateName = (name: string): boolean => {
  return name.trim().length >= 2 && /^[a-zA-Zа-яА-ЯёЁ\s-]+$/.test(name);
};

export const validateTelegramBotToken = (token: string): { isValid: boolean; message?: string } => {
  if (!token.trim()) {
    return { isValid: false, message: 'Telegram Bot Token is required' };
  }

  // Telegram bot tokens follow the format: 123456789:ABCdefGHIjklmnoPQRstuVWXyz
  const tokenRegex = /^\d+:[A-Za-z0-9_-]{35}$/;
  
  if (!tokenRegex.test(token)) {
    return { 
      isValid: false, 
      message: 'Invalid token format. Expected format: 123456789:ABCdefGHIjklmnoPQRstuVWXyz' 
    };
  }

  return { isValid: true };
};
