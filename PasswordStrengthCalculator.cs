namespace PasswordGenerator
{
    static class PasswordStrengthCalculator
    {
        private static string WeakPassword = "Weak";
        private static string ModeratePassword = "Moderate";
        private static string StrongPassword = "Strong";

        static string PasswordStrength(
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
