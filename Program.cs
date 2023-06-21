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
    }
}
