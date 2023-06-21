namespace PasswordGenerator
{
    static class PasswordStrengthCalculator
    {
        static string WeakPassword = "Weak";
        static string ModeratePassword = "Moderate";
        static string StrongPassword = "Strong";

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
