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
