namespace PasswordGenerator
{
    static class PasswordStrengthCalculator
    {
        private static readonly string WeakPassword = "Weak";
        private static readonly string ModeratePassword = "Moderate";
        private static readonly string StrongPassword = "Strong";

        public static string CalculatePasswordStrength(
            int length,
            bool useUpperCase,
            bool useLowerCase,
            bool useSpecialChars,
            bool useNumbers
        )
        {
            if (length < 8 || (!useUpperCase || !useLowerCase || !useNumbers || !useSpecialChars))
                return WeakPassword;
            else if (
                length < 12 || (!useUpperCase && !useLowerCase) || (!useNumbers && !useSpecialChars)
            )
                return ModeratePassword;
            else
                return StrongPassword;
        }
    }
}
