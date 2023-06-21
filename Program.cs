using System.Text;

namespace PasswordGenerator
{
    class Program
    {
        static void Main(string[] args)
        {
            int passwordLength = GetPasswordLength();
            bool useUpperCase = GetBooleanInput("Use uppercase letters? (Y/N): ");
            bool useLowerCase = GetBooleanInput("Use lowercase letters? (Y/N): ");
            bool useSpecialChars = GetBooleanInput("Use special characters? (Y/N): ");
            bool useNumbers = GetBooleanInput("Use numbers? (Y/N): ");

            if (!useUpperCase && !useLowerCase && !useSpecialChars && !useNumbers)
            {
                LogMessage(
                    "You must select at least one character option. Please try again.",
                    ConsoleColor.DarkRed
                );
                return;
            }

            string generatedPassword = GeneratePassword(
                passwordLength,
                useUpperCase,
                useLowerCase,
                useSpecialChars,
                useNumbers
            );

            string strength = PasswordStrengthCalculator.CalculatePasswordStrength(
                passwordLength,
                useUpperCase,
                useLowerCase,
                useSpecialChars,
                useNumbers
            );

            DisplayPasswordWithStrength(generatedPassword, strength);

            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }

        static int GetPasswordLength()
        {
            int passwordLength;
            while (true)
            {
                Console.Write("Enter the length of the password: ");
                if (int.TryParse(Console.ReadLine(), out passwordLength) && passwordLength > 0)
                {
                    break;
                }
                Console.WriteLine("Invalid input. Please enter a positive integer.");
            }
            return passwordLength;
        }

        static bool GetBooleanInput(string message)
        {
            while (true)
            {
                Console.Write(message);
                string? input = Console.ReadLine()?.ToUpper();
                if (input == "Y" || input == "YES")
                {
                    return true;
                }
                else if (input == "N" || input == "NO")
                {
                    return false;
                }
                Console.WriteLine("Invalid input. Please enter 'Y' or 'N'.");
            }
        }

        static string GeneratePassword(
            int length,
            bool useUpperCase,
            bool useLowerCase,
            bool useSpecialChars,
            bool useNumbers
        )
        {
            string uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
            string lowercaseChars = "abcdefghijklmnopqrstuvwxyz";
            string specialChars = "!@#$%^&*()";
            string numberChars = "0123456789";

            StringBuilder password = new StringBuilder();
            Random random = new Random();

            StringBuilder charPool = new StringBuilder();
            if (useUpperCase)
                charPool.Append(uppercaseChars);
            if (useLowerCase)
                charPool.Append(lowercaseChars);
            if (useSpecialChars)
                charPool.Append(specialChars);
            if (useNumbers)
                charPool.Append(numberChars);

            for (int i = 0; i < length; i++)
            {
                password.Append(charPool[random.Next(0, charPool.Length)]);
            }

            return password.ToString();
        }

        static void DisplayPasswordWithStrength(string password, string strength)
        {
            Console.WriteLine("Generated Password: " + password);
            Console.Write("Password Strength: ");

            switch (strength)
            {
                case "Weak":
                    LogMessage(strength, ConsoleColor.Red);
                    break;
                case "Moderate":
                    LogMessage(strength, ConsoleColor.Yellow);
                    break;
                case "Strong":
                    LogMessage(strength, ConsoleColor.Green);
                    break;
                default:
                    Console.WriteLine(strength);
                    break;
            }
        }

        static void LogMessage(string message, ConsoleColor color)
        {
            Console.ForegroundColor = color;
            Console.WriteLine(message);
            Console.ResetColor();
        }
    }
}
