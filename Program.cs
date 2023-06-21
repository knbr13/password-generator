using System.Text;

namespace PasswordGenerator
{
    class Program
    {
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
                if (message == "Y" || message == "YES")
                {
                    return true;
                }
                else if (message == "N" || message == "NO")
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
    }
}
