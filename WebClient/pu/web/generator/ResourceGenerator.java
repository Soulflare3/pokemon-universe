package pu.web.generator;

import java.io.File;
import java.io.FileFilter;
import java.io.PrintWriter;
import java.util.HashSet;
import java.util.Set;

import com.google.gwt.core.client.GWT;
import com.google.gwt.core.ext.Generator;
import com.google.gwt.core.ext.GeneratorContext;
import com.google.gwt.core.ext.TreeLogger;
import com.google.gwt.core.ext.UnableToCompleteException;
import com.google.gwt.core.ext.typeinfo.JClassType;
import com.google.gwt.core.ext.typeinfo.NotFoundException;
import com.google.gwt.core.ext.typeinfo.TypeOracle;
import com.google.gwt.resources.client.ClientBundleWithLookup;
import com.google.gwt.resources.client.DataResource;
import com.google.gwt.resources.client.ImageResource;
import com.google.gwt.resources.client.ResourcePrototype;
import com.google.gwt.resources.client.TextResource;
import com.google.gwt.user.rebind.ClassSourceFileComposerFactory;
import com.google.gwt.user.rebind.SourceWriter;

public class ResourceGenerator extends Generator
{
	private static FileFilter fileFilter = new FileFilter()
	{
		@Override
		public boolean accept(File file)
		{
			if (file.isDirectory())
			{
				return true;
			} 
			else
			{
				String extension = getExtension(file.getName());
				if(extension.equals(".png"))
					return true;
			}
			return false;
		}
	};

	private static String getContentType(TreeLogger logger, File file)
	{
		String name = file.getName().toLowerCase();
		int pos = name.lastIndexOf('.');
		String extension = pos == -1 ? "" : name.substring(pos);
		
		String contentType = "";
		if(extension.equals(".png"))
			contentType = "image/png";
		else
			contentType = "application/octet-stream";
			
		return contentType;
	}

	private static boolean isValidMethodName(String methodName)
	{
		return methodName.matches("^[a-zA-Z_$][a-zA-Z0-9_$]*$");
	}

	private static String stripExtension(String filename)
	{
		return filename.replaceFirst("\\.[^.]+$", "");
	}

	private static String getExtension(String filename)
	{
		return filename.replaceFirst(".*(\\.[^.]+)$", "$1");
	}

	@Override
	public String generate(TreeLogger logger, GeneratorContext context, String typeName) throws UnableToCompleteException
	{
		TypeOracle typeOracle = context.getTypeOracle();
		
		JClassType userType;
		try
		{
			userType = typeOracle.getType(typeName);
		} 
		catch (NotFoundException e)
		{
			logger.log(TreeLogger.ERROR, "Unable to find metadata for type: " + typeName, e);
			throw new UnableToCompleteException();
		}
		
		String packageName = userType.getPackage().getName();
		String className = userType.getName();
		className = className.replace('.', '_');

		if (userType.isInterface() == null)
		{
			logger.log(TreeLogger.ERROR, userType.getQualifiedSourceName() + " is not an interface", null);
			throw new UnableToCompleteException();
		}

		ClassSourceFileComposerFactory composerFactory = new ClassSourceFileComposerFactory(packageName, className + "Impl");
		composerFactory.addImplementedInterface(userType.getQualifiedSourceName());

		composerFactory.addImport(ClientBundleWithLookup.class.getName());
		composerFactory.addImport(DataResource.class.getName());
		composerFactory.addImport(GWT.class.getName());
		composerFactory.addImport(ImageResource.class.getName());
		composerFactory.addImport(ResourcePrototype.class.getName());
		composerFactory.addImport(TextResource.class.getName());

		File classesDirectory = new File("war/WEB-INF/classes/");

		File resourcesDirectory = new File(classesDirectory, packageName.replace('.', '/'));
		
		String baseClassesPath = classesDirectory.getPath();

		Set<File> files = getFiles(resourcesDirectory, fileFilter);
		Set<String> methodNames = new HashSet<String>();
		
		PrintWriter pw = context.tryCreate(logger, packageName, className + "Impl");
		if (pw != null)
		{
			SourceWriter sw = composerFactory.createSourceWriter(context, pw);

			sw.println("public ResourcePrototype[] getResources() {");
			sw.indent();
			sw.println("return MyBundle.INSTANCE.getResources();");
			sw.outdent();
			sw.println("}");

			sw.println("public ResourcePrototype getResource(String name) {");
			sw.indent();
			sw.println("return MyBundle.INSTANCE.getResource(name);");
			sw.outdent();
			sw.println("}");

			sw.println("static interface MyBundle extends ClientBundleWithLookup {");
			sw.indent();
			sw.println("MyBundle INSTANCE = GWT.create(MyBundle.class);");

			for (File file : files)
			{
				String filepath = file.getPath();
				String relativePath = filepath.replace(baseClassesPath, "").replace('\\', '/').replaceFirst("^/", "");
				String filename = file.getName();
				String contentType = getContentType(logger, file);
				String methodName = stripExtension(filename);

				if (!isValidMethodName(methodName))
				{
					//logger.log(TreeLogger.WARN, "Skipping invalid method name (" + methodName + ") due to: " + relativePath);
					//continue;
					methodName = "res_" + methodName;
				}
				if (!methodNames.add(methodName))
				{
					logger.log(TreeLogger.WARN, "Skipping duplicate method name due to: " + relativePath);
					continue;
				}

				Class<? extends ResourcePrototype> returnType = getResourcePrototype(contentType);

				// generate method
				sw.println();
				sw.println("@Source(\"" + relativePath + "\")");
				sw.println(returnType.getName() + " " + methodName + "();");
			}

			sw.outdent();
			sw.println("}");

			sw.commit(logger);
		}
		logger.log(TreeLogger.INFO, "herp", null);
		return composerFactory.getCreatedClassName();
	}

	private HashSet<File> getFiles(File dir, FileFilter filter)
	{
		HashSet<File> fileList = new HashSet<File>();
		File[] files = dir.listFiles(filter);
		for (int i = 0; i < files.length; i++)
		{
			File f = files[i];
			if (f.isFile())
			{
				fileList.add(f);
			} 
			else
			{
				if (filter.accept(f))
				{
					fileList.addAll(getFiles(f, filter));
				}
			}
		}
		return fileList;
	}

	private Class<? extends ResourcePrototype> getResourcePrototype(String contentType)
	{
		Class<? extends ResourcePrototype> returnType;
		if (contentType.startsWith("image/"))
		{
			returnType = ImageResource.class;
		} 
		else if (contentType.startsWith("text/"))
		{
			returnType = TextResource.class;
		} 
		else
		{
			returnType = DataResource.class;
		}
		return returnType;
	}
}
